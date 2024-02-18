/* eslint-disable @next/next/no-img-element  */
/* eslint-disable jsx-ally/alt-text */
// @ts-nocheck
import { ImageResponse } from 'next/server'

export const runtime = 'edge';


export async function GET(req: Request) {
  const imageData = await fetch(new URL('../../public/base.png', import.meta.url)).then(
    (res) => res.arrayBuffer(),
  );
  return new ImageResponse (
    <div>
       <img src={imageData} width={300} height={217} />
       <div> Stake </div>
    </div> 
    )
  }