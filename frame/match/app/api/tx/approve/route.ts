import { FrameRequest, getFrameMessage } from '@coinbase/onchainkit';
import { NextRequest, NextResponse } from 'next/server';
import { encodeFunctionData, parseEther, parseUnits } from 'viem';
import { base } from 'viem/chains';
import PrizePool from "../contracts/PrizePool.json";
import IERC20 from "../contracts/IERC20.json";
import type { FrameTransactionResponse } from '@coinbase/onchainkit/src/frame'

async function getResponse(req: NextRequest): Promise<NextResponse> {
  let prizePoolAddr = '0xd9D454387F1cF48DB5b7D40C5De9d5bD9a92C1F8';
  
  const body: FrameRequest = await req.json();
  const searchParams = req.nextUrl.searchParams;
  const gameId:any = searchParams.get("gameId");
  const gameName:any = searchParams.get("gameName");
  const gameSetup:any = searchParams.get("gameSetup");
  const stakeAmount:any = searchParams.get("stakeAmount");
  const creatorFid:any = searchParams.get("creatorFid");
  
  let queryParams = `gameId=${gameId}&&gameName=${gameName}&&gameSetup=${gameSetup}&&stakeAmount=${stakeAmount}&&creatorFid=${creatorFid}`
  let baseUrl = "https://wag3r-bot-gamma.vercel.app/?" + queryParams;
  const { isValid } = await getFrameMessage(body, { neynarApiKey: 'NEYNAR_ONCHAIN_KIT' });


  const buttonId = body.untrustedData.buttonIndex;
  if (buttonId == 1) return NextResponse.redirect(baseUrl);

  if (!isValid) {
    return new NextResponse('Message not valid', { status: 500 });
  }

  if (buttonId == 2) {
    return handlesStake2StartMatch()
  }

  return NextResponse.redirect(baseUrl)
}

export async function POST(req: NextRequest): Promise<Response> {
  return getResponse(req);
}

export const dynamic = 'force-dynamic';


function handlesStake2StartMatch() {
  let prizePoolAddr = '0x5E0f293eEBa536e6cDeB0B9da03d1b5335dC29De'; // contract yet to be deployed 
  const data = encodeFunctionData({
    abi: IERC20.abi,
    functionName: 'approve',
    args: [prizePoolAddr, parseUnits("0.1", 6)],
  });

  const txData: FrameTransactionResponse = {
    chainId: `eip155:${base.id}`,
    method: 'eth_sendTransaction',
    params: {
      abi: [],
      data,
      to: '0x833589fCD6eDb6E08f4c7C32D4f71b54bdA02913',
      value: '0x0', // 0.01 ETH
    },
  };
  return NextResponse.json(txData);
}



