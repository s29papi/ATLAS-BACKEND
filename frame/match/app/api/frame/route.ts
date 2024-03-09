import { FrameRequest, getFrameMessage } from '@coinbase/onchainkit';
import { NextRequest, NextResponse } from 'next/server';

async function getResponse(req: NextRequest): Promise<NextResponse> {
  const body: FrameRequest = await req.json();
  const searchParams = req.nextUrl.searchParams;
  const gameId:any = searchParams.get("gameId");
  const gameName:any = searchParams.get("gameName");
  const gameSetup:any = searchParams.get("gameSetup");
  const stakeAmount:any = searchParams.get("stakeAmount");
  const creatorFid:any = searchParams.get("creatorFid");
  // const buttonId = body.untrustedData.buttonIndex;
  
  let queryParams = `gameId=${gameId}&&gameName=${gameName}&&gameSetup=${gameSetup}&&stakeAmount=${stakeAmount}&&creatorFid=${creatorFid}`
  let button2ImageUrl = "https://wag3r-bot-gamma.vercel.app/og/approve?" + `${queryParams}`
  let postUrl = "https://wag3r-bot-gamma.vercel.app/api/frame/stake?" + `${queryParams}`

    return new NextResponse(`<!DOCTYPE html><html><head>
          <title>Start My Match</title>
          <meta property="fc:frame" content="vNext" />        
          <meta property="fc:frame:image" content="${button2ImageUrl}"/>
          <meta property="fc:frame:button:1" content="Back" />
          <meta property="fc:frame:button:1:action" content="post"/>
          <meta property="fc:frame:button:2" content="Approve" />
          <meta property="fc:frame:button:2:action" content="tx"/>
          <meta property="fc:frame:button:2:target" content="https://wag3r-bot-gamma.vercel.app/api/frame/stake"/>
          <meta property="fc:frame:post_url" content="${postUrl}"/>
      </head></html>`);
}

export async function POST(req: NextRequest): Promise<Response> {
  return getResponse(req);
}

export const dynamic = 'force-dynamic';
