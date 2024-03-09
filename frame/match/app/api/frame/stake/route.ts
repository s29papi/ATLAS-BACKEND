// import { FrameRequest, getFrameMessage } from '@coinbase/onchainkit';
import { FrameRequest, getFrameMessage } from '@coinbase/onchainkit';
import { NextRequest, NextResponse } from 'next/server';
import { encodeFunctionData, parseEther } from 'viem';
import { base } from 'viem/chains';
import PrizePool from "./contracts/PrizePool.json"
import IERC20 from "./contracts/IERC20.json"
import type { FrameTransactionResponse } from '@coinbase/onchainkit/src/frame'

async function getResponse(req: NextRequest): Promise<NextResponse> {
  let prizePoolAddr = '0xd9D454387F1cF48DB5b7D40C5De9d5bD9a92C1F8';
  let vusdcAddr = '0x4dd745f5aca5b63999cb097c0c11cc4338e2febf';
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
  // let vusdcAddr = '0x4dd745f5aca5b63999cb097c0c11cc4338e2febf';
  let prizePoolAddr = '0xd9D454387F1cF48DB5b7D40C5De9d5bD9a92C1F8';
    const data = encodeFunctionData({
      abi: IERC20.abi,
      functionName: 'approve',
      args: [prizePoolAddr, 1000000],
    });

    const txData: FrameTransactionResponse = {
      chainId: `eip155:${base.id}`,
      method: 'eth_sendTransaction',
      params: {
        abi: [],
        data,
        to: '0x4dd745f5aca5b63999cb097c0c11cc4338e2febf',
        value: parseEther('0.00').toString(), // 0.01 ETH
      },
    };
    return NextResponse.json(txData);
            //     return new NextResponse(`<!DOCTYPE html><html><head>
            //     <title>Start My Match</title>
            //     <meta property="fc:frame" content="vNext" />
            //     <meta property="fc:frame:image" content="https://wag3r-bot.vercel.app/A-New-Challenger-Has-Entered-The-Ring-Resize.png"/>
            //     <meta property="fc:frame:button:1" content="Start My Match" />
            //     <meta property="fc:frame:button:1:action" content="post"/>
            //     <meta property="fc:frame:post_url" content="https://wag3r-bot.vercel.app/api/frame/start-match"/>
            // </head></html>`);
}




// wager: is the second buttonId
//   const searchParams = req.nextUrl.searchParams
//   const gameId:any = searchParams.get("game-id")