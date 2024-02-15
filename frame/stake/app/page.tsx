// spec: https://docs.farcaster.xyz/reference/frames/spec
// playground: https://og-playground.vercel.app/

import { getFrameMetadata } from '@coinbase/onchainkit';
import type { Metadata, ResolvingMetadata  } from 'next';

type Props = {
  params: { id: string }
  searchParams: { [key: string]: string | string[] | undefined }
}



let frameMetadata;

let postUrl;

export async function generateMetadata(
  { params, searchParams }: Props,
  parent: ResolvingMetadata
){
  postUrl = 'https://wag3r-bot-stake.vercel.app/api'; 
}


frameMetadata = getFrameMetadata({
  buttons: [
      {label: 'Accept Challenge & Stake Tokens', action: 'post_redirect'},
  ],
  image: 'https://wag3r-bot-stake.vercel.app/stadium-figma-test-1.png',
  post_url: postUrl,
});

if (!frameMetadata) {
  throw new Error('Project ID is not defined')
}

export const metadata: Metadata = {
  title: 'Refuel-Frame by socket.',
  description: 'Follow this user, Like the post, and Refuel.',
  openGraph: {
    title: 'Refuel-Frame by socket.',
    description: 'Follow this user, Like the post, and Refuel.',
    images: [`https://wag3r-bot-stake.vercel.app/stadium-figma-test-1.png`],
  },
  other: {
    ...frameMetadata,
  },
};


export default function Page({
  params,
  searchParams,
}: {
  params: { slug: string };
  searchParams?: { [key: string]: string | string[] | undefined };
}) {
  
  return <h1>{searchParams?.gameid || "Hello!"}</h1>;
}



