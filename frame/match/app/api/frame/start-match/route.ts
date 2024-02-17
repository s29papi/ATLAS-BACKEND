import { FrameRequest, getFrameMessage } from '@coinbase/onchainkit';
import { NextRequest, NextResponse } from 'next/server';

async function getResponse(req: NextRequest): Promise<NextResponse> {
  const body: FrameRequest = await req.json();
  

  return new NextResponse(`<!DOCTYPE html><html><head>
            <title>Stream Link</title>
            <meta property="fc:frame" content="vNext" />
            <meta property="fc:frame:image" content="https://wag3r-bot.vercel.app/You-Vs-Me-Rescale.png"/>
            <meta property="fc:frame:input:text" content="Insert Stream Link"/>
            <meta property="fc:frame:button:1" content="No Stream" />
            <meta property="fc:frame:button:1:action" content="post"/>
            <meta property="fc:frame:button:2" content="Continue" />
            <meta property="fc:frame:button:2:action" content="post"/>
            <meta property="fc:frame:button:3" content="View Match Details" />
            <meta property="fc:frame:button:3:action" content="post"/>
            <meta property="fc:frame:post_url" content="https://wag3r-bot.vercel.app/api/frame/ready-up"/>
            </head></html>`);
}

export async function POST(req: NextRequest): Promise<Response> {
  return getResponse(req);
}

export const dynamic = 'force-dynamic';






// wager: is the second buttonId
//   const searchParams = req.nextUrl.searchParams
//   const gameId:any = searchParams.get("game-id")