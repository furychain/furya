package keeper_test

import (
	fmt "fmt"

	sharedtypes "github.com/furychain/furya/shared/types"
	"github.com/furychain/furya/testutil/sample"
	"github.com/furychain/furya/x/rollapp/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (suite *RollappTestSuite) createRollappAndVerify(numOfAddresses int, expectedErr error, rollappsExpect *[]*types.Rollapp) {
	goCtx := sdk.WrapSDKContext(suite.ctx)
	// generate sequences address
	addresses := sample.GenerateAddresses(numOfAddresses)
	// rollapp is the rollapp to create
	rollapp := types.MsgCreateRollapp{
		Creator:               alice,
		RollappId:             fmt.Sprintf("%s%d", "rollapp", numOfAddresses),
		CodeStamp:             "",
		GenesisPath:           "",
		MaxWithholdingBlocks:  1,
		MaxSequencers:         1,
		PermissionedAddresses: sharedtypes.Sequencers{Addresses: addresses},
	}
	// rollappExpect is the expected result of creating rollapp
	rollappExpect := types.Rollapp{
		RollappId:             rollapp.GetRollappId(),
		Creator:               rollapp.GetCreator(),
		Version:               0,
		CodeStamp:             rollapp.GetCodeStamp(),
		GenesisPath:           rollapp.GetGenesisPath(),
		MaxWithholdingBlocks:  rollapp.GetMaxWithholdingBlocks(),
		MaxSequencers:         rollapp.GetMaxSequencers(),
		PermissionedAddresses: rollapp.GetPermissionedAddresses(),
	}
	// create rollapp
	createResponse, err := suite.msgServer.CreateRollapp(goCtx, &rollapp)
	if expectedErr != nil {
		suite.EqualError(err, expectedErr.Error())
		return
	}
	suite.Require().Nil(err)
	suite.Require().EqualValues(types.MsgCreateRollappResponse{}, *createResponse)

	// query the specific rollapp
	queryResponse, err := suite.queryClient.Rollapp(goCtx, &types.QueryGetRollappRequest{
		RollappId: rollapp.GetRollappId(),
	})
	if queryResponse.Rollapp.PermissionedAddresses.Addresses == nil {
		queryResponse.Rollapp.PermissionedAddresses.Addresses = []string{}
	}
	suite.Require().Nil(err)
	suite.Require().EqualValues(&rollappExpect, &queryResponse.Rollapp)

	// add the rollapp to the list of get all expected list
	newRollappsExpect := *rollappsExpect
	newRollappsExpect = append(newRollappsExpect, &rollappExpect)
	// verify that query all contains all the rollapps that were created
	rollappsRes, totalRes := getAll(suite)
	suite.Require().EqualValues(totalRes, numOfAddresses+1)
	vereifyAll(suite, newRollappsExpect, rollappsRes)
	*rollappsExpect = newRollappsExpect
}

func (suite *RollappTestSuite) createRollappFromWhitelist(expectedErr error, deployerWhitelist []types.DeployerParams) {
	suite.SetupTest(deployerWhitelist...)

	// rollappsExpect is the expected result of query all
	rollappsExpect := []*types.Rollapp{}

	// test 10 rollap creations
	for i := 0; i < 10; i++ {
		suite.createRollappAndVerify(i, expectedErr, &rollappsExpect)
	}

}

func (suite *RollappTestSuite) TestCreateRollapp() {
	suite.createRollappFromWhitelist(nil, nil)
}

func (suite *RollappTestSuite) TestCreateRollappFromWhitelist() {
	suite.createRollappFromWhitelist(nil, []types.DeployerParams{{Address: alice, MaxRollapps: 0}})
}

func (suite *RollappTestSuite) TestCreateRollappUnauthorizedRollappCreator() {
	suite.createRollappFromWhitelist(types.ErrUnauthorizedRollappCreator, []types.DeployerParams{{Address: bob, MaxRollapps: 0}})
}

func (suite *RollappTestSuite) TestCreateRollappAlreadyExists() {
	suite.SetupTest()
	goCtx := sdk.WrapSDKContext(suite.ctx)

	// rollapp is the rollapp to create
	rollapp := types.MsgCreateRollapp{
		Creator:               alice,
		RollappId:             "rollapp1",
		CodeStamp:             "",
		GenesisPath:           "",
		MaxWithholdingBlocks:  1,
		MaxSequencers:         1,
		PermissionedAddresses: sharedtypes.Sequencers{},
	}
	_, err := suite.msgServer.CreateRollapp(goCtx, &rollapp)
	suite.Require().Nil(err)

	_, err = suite.msgServer.CreateRollapp(goCtx, &rollapp)
	suite.EqualError(err, types.ErrRollappExists.Error())
}

func (suite *RollappTestSuite) TestCreateRollappExceedMaxRollapps() {
	deployerWhitelist := []types.DeployerParams{{Address: alice, MaxRollapps: 10}}

	suite.SetupTest(deployerWhitelist...)

	// rollappsExpect is the expected result of query all
	rollappsExpect := []*types.Rollapp{}

	// test 10 rollap creations
	for i := 0; i < 10; i++ {
		suite.createRollappAndVerify(i, nil, &rollappsExpect)
	}

	suite.createRollappAndVerify(10, types.ErrRollappCreatorExceedMaximumRollapps, &rollappsExpect)
}
