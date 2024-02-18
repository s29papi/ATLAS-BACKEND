/* eslint-disable @next/next/no-img-element  */
/* eslint-disable jsx-ally/alt-text */
// @ts-nocheck
import { ImageResponse } from 'next/server'

import { Image } from 'next/image'

import base from '../../public/base.png'
// App router includes @vercel/og.
// No need to install it.
// import base from '../../public/base.png'
export const runtime = 'edge';


export async function GET() {
  const imageData = await fetch(new URL('../../public/base.png', import.meta.url)).then(
    (res) => res.arrayBuffer(),
  );


  console.log(imageData)
    return new ImageResponse(
      (
        <div
          // style={{
          //   width: "100vw",
          //   height: "100vh",
          //   display: 'flex',
          //   alignItems: 'center',
          //   justifyContent: 'center',
          //   fontSize: 128,
          //   // background: 'lavender',
          //   // backgroundImage: `url(${}),`
          //   backgroundPosition: "center",
          //   backgroundSize: "cover",
          //   backgroundRepeat: "no-repeat",
          // }}
        >
          <Image alt='Base Stake Page' src={base} style={{maxWidth: '100%', height: 'auto'}}/>
          {/* Bizzy */}
          {/* <img width="800" height="419" src={imageData}  /> */}
        </div>
      ),
      {
        width: 800,
        height: 419,
      }
    )
  }