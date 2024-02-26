import { getFrameMetadata } from '@coinbase/onchainkit';
import type { Metadata, ResolvingMetadata } from 'next'

type Props = {
  params: { gameId: string }
  searchParams: { [key: string]: string | string[] | undefined }
}
 

export async function generateMetadata(
  { params, searchParams }: Props,
  parent: ResolvingMetadata
): Promise<Metadata> { 
  const gameid = searchParams["gameId"];
  // work on this
  let postUrl = "https://wag3r-bot.vercel.app/api/frame?gameId=" + `${gameid}`;

  const frameMetadata = getFrameMetadata({
    buttons: [
        {label: 'Accept Challenge', action: 'post'},
        {label: 'Acount', action: 'post'},
    ],
    image: 'https://wag3r-bot-gamma.vercel.app/og/landingframe',
    post_url: postUrl,
  });

  return {
    title: 'Match By Versus.',
    description: 'Frontend Match Management for Versus App.',
    openGraph: {
      title: 'Match By Versus.',
      description: 'Frontend Match Management for Versus App.',
      images: [`https://wag3r-bot-gamma.vercel.app/og/landingframe`],
    },
    other: {
      ...frameMetadata,
    }
  }
}

// A-New-Challenger-Has-Entered-The-Ring-Resize


export default function Page({
  params,
  searchParams,
}: {
  params: { slug: string };
  searchParams?: { [key: string]: string | string[] | undefined };
}) {
  return <h1>Match By Versus iiiooo.</h1>;
}



// spec: https://docs.farcaster.xyz/reference/frames/spec  
// playground: https://og-playground.vercel.app/


