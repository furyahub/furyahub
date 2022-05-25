package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/mining/types"
)

func (k msgServer) AddRewardPool(goCtx context.Context, msg *types.MsgAddRewardPool) (*types.MsgAddRewardPoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	user, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}

	stakePool, found := k.Keeper.GetStakePool(ctx, msg.StakePoolIndex)
	if !found {
		return nil, types.ErrStakePoolNotExist
	}

	maxRewardPoolNumber := k.Keeper.GetMaxRewardPoolNumber(ctx)
	if len(stakePool.RewardPools) >= int(maxRewardPoolNumber) {
		return nil, types.ErrRewardPoolNumberReachLimit
	}

	rewardToken, found := k.Keeper.GetRewardToken(ctx, msg.RewardTokenDenom)
	if !found {
		return nil, types.ErrRewardTokenNotSupport
	}
	if msg.TotalRewardAmount.LT(rewardToken.MinTotalRewardAmount) {
		return nil, types.ErrTotalRewardAmountLessThanLimit
	}

	curBlockTime := uint64(ctx.BlockTime().Unix())

	willUseIndex := k.Keeper.GetRewardPoolNextIndex(ctx, msg.StakePoolIndex)
	willUseLastRewardTimestamp := msg.StartTimestamp
	if msg.StartTimestamp < curBlockTime {
		willUseLastRewardTimestamp = curBlockTime
	}

	rewardTokens := sdk.NewCoins(sdk.NewCoin(msg.RewardTokenDenom, msg.TotalRewardAmount))
	if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, user, types.ModuleName, rewardTokens); err != nil {
		return nil, err
	}

	stakePool.RewardPools = append(stakePool.RewardPools, &types.RewardPool{
		Index:               willUseIndex,
		RewardTokenDenom:    msg.RewardTokenDenom,
		TotalRewardAmount:   msg.TotalRewardAmount,
		LeftRewardAmount:    msg.TotalRewardAmount,
		RewardPerSecond:     msg.RewardPerSecond,
		StartTimestamp:      msg.StartTimestamp,
		RewardPerPower:      sdk.ZeroInt(),
		LastRewardTimestamp: willUseLastRewardTimestamp,
	})

	k.Keeper.SetRewardPoolIndex(ctx, msg.StakePoolIndex, willUseIndex)
	k.Keeper.SetStakePool(ctx, stakePool)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeAddRewardPool,
			sdk.NewAttribute(types.AttributeKeyAccount, msg.Creator),
			sdk.NewAttribute(types.AttributeKeyStakePoolIndex, fmt.Sprintf("%d", msg.StakePoolIndex)),
			sdk.NewAttribute(types.AttributeKeyRewardTokenDenom, msg.RewardTokenDenom),
			sdk.NewAttribute(types.AttributeKeyTotalRewardAmount, msg.TotalRewardAmount.String()),
			sdk.NewAttribute(types.AttributeKeyRewardPerSecond, msg.RewardPerSecond.String()),
			sdk.NewAttribute(types.AttributeKeyStartTimestamp, fmt.Sprintf("%d", msg.StartTimestamp)),
			sdk.NewAttribute(types.AttributeKeyLastRewardTimestamp, fmt.Sprintf("%d", willUseLastRewardTimestamp)),
		),
	)

	return &types.MsgAddRewardPoolResponse{}, nil
}