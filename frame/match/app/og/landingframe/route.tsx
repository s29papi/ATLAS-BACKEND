/* eslint-disable @next/next/no-img-element  */
/* eslint-disable jsx-ally/alt-text */
// @ts-nocheck
import { ImageResponse } from 'next/server'

export const runtime = 'edge';


export async function GET(req: Request) {
    const imageData = await fetch(new URL('../../../public/You-Vs-Me-Rescale.png', import.meta.url)).then(
        (res) => res.arrayBuffer(),
      );
    return new ImageResponse( 
        (
            <div style={{position: 'relative', display: 'flex'}}>
                 <img 
                        src={imageData}
                    />
            </div>
        )
    )
}

