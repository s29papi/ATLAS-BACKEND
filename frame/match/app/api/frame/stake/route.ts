import { FrameRequest, getFrameMessage } from '@coinbase/onchainkit';
import { NextRequest, NextResponse } from 'next/server';

async function getResponse(req: NextRequest): Promise<NextResponse> {
  const body: FrameRequest = await req.json();
  const buttonId = body.untrustedData.buttonIndex;

// wager: is the second buttonId
//   const searchParams = req.nextUrl.searchParams
//   const gameId:any = searchParams.get("game-id")

  if (buttonId == 2) {
    return new NextResponse(`<!DOCTYPE html><html><head>
          <title>Start My Match</title>
          <meta property="fc:frame" content="vNext" />
          <meta property="fc:frame:image" content="https://wag3r-bot.vercel.app/A-New-Challenger-Has-Entered-The-Ring-Resize.png"/>
          <meta property="fc:frame:button:1" content="Start My Match" />
          <meta property="fc:frame:button:1:action" content="post"/>
          <meta property="fc:frame:post_url" content=""/>
      </head></html>`);
  }

  return NextResponse.redirect("https://wag3r-bot.vercel.app/")
}

export async function POST(req: NextRequest): Promise<Response> {
  return getResponse(req);
}

export const dynamic = 'force-dynamic';
