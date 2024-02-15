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
  let postUrl = "https://wag3r-bot.vercel.app/api?gameid=" + `${gameid}`;

  const frameMetadata = getFrameMetadata({
    buttons: [
        {label: 'Accept Challenge & Stake Tokens', action: 'post_redirect'},
    ],
    image: 'https://wag3r-bot-stake.vercel.app/stadium-figma-test-1.png',
    post_url: postUrl,
  });

  return {
    title: 'Stake By Versus.',
    description: 'Frontend Stake Management for Versus App.',
    openGraph: {
      title: 'Stake By Versus.',
      description: 'Frontend Stake Management for Versus App.',
      images: [`https://wag3r-bot-stake.vercel.app/stadium-figma-test-1.png`],
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
  return <h1>Stake By Versus.</h1>;
}


// ref: https://www.pinata.cloud/blog/how-to-make-a-frame-on-farcaster-using-ipfs
// spec: https://docs.farcaster.xyz/reference/frames/spec
// playground: https://og-playground.vercel.app/

