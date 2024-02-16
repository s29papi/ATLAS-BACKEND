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
  const gameid = searchParams["gameid"];
  let postUrl = "https://wag3r-bot-stake.vercel.app/api?gameid=" + `${gameid}`;

  const frameMetadata = getFrameMetadata({
    buttons: [
        {label: 'View Txn', action: 'post_redirect'},
        {label: 'Continue', action: 'post'},
    ],
    image: 'https://wag3r-bot.vercel.app/stadium-figma-test-2b.png',
    post_url: postUrl,
  });

  return {
    title: 'Match By Versus.',
    description: 'Frontend Match Management for Versus App.',
    openGraph: {
      title: 'Match By Versus.',
      description: 'Frontend Match Management for Versus App.',
      images: [`https://wag3r-bot.vercel.app/stadium-figma-test-2b.png`],
    },
    other: {
      ...frameMetadata,
    }
  }
}


export default function Page({
  params,
  searchParams,
}: {
  params: { slug: string };
  searchParams?: { [key: string]: string | string[] | undefined };
}) {
  return <h1>Match By Versus.</h1>;
}



// spec: https://docs.farcaster.xyz/reference/frames/spec
// playground: https://og-playground.vercel.app/

