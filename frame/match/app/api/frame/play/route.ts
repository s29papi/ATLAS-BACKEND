import { FrameRequest, getFrameMessage } from '@coinbase/onchainkit';
import { NextRequest, NextResponse } from 'next/server';

async function getResponse(req: NextRequest): Promise<NextResponse> {
  const body: FrameRequest = await req.json();
  

  return new NextResponse(`<!DOCTYPE html><html><head>
            <title>Stream Link</title>
            <meta property="fc:frame" content="vNext" />
            <meta property="fc:frame:image" content="https://wag3r-bot.vercel.app/You-Vs-Me-Rescale.png"/>
            <meta property="fc:frame:button:1" content="You are now ready to play!" />
            <meta property="fc:frame:button:1:action" content="post"/>
           
            </head></html>`);
}

export async function POST(req: NextRequest): Promise<Response> {
  return getResponse(req);
}

export const dynamic = 'force-dynamic';



// <meta property="fc:frame:post_url" content="https://wag3r-bot.vercel.app/"/>