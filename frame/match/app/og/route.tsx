/* eslint-disable @next/next/no-img-element  */
/* eslint-disable jsx-ally/alt-text */
// @ts-nocheck
import { ImageResponse } from 'next/server'
// App router includes @vercel/og.
// No need to install it.
// import base from '../../public/base.png'
export const runtime = 'edge';

export async function GET() {
  const imageData = await fetch(new URL('./base.png', import.meta.url)).then(
    (res) => res.arrayBuffer(),
  );


  console.log(imageData)
    return new ImageResponse(
      (
        <div
          style={{
            width: '100%',
            height: '100%',
            display: 'flex',
            alignItems: 'center',
            justifyContent: 'center',
            fontSize: 128,
            background: 'lavender',
            // backgroundImage: `url(${}),`
          }}
        >
          <img width="256" height="256" src={imageData} />
        </div>
      )
    )
  }