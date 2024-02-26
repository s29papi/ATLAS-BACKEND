/* eslint-disable @next/next/no-img-element  */
/* eslint-disable jsx-ally/alt-text */
// @ts-nocheck
import { ImageResponse } from 'next/server'

export const runtime = 'edge';


export async function GET(req: Request) {
    const imageData = await fetch(new URL('../../../public/You-Vs-Me-Rescale.png', import.meta.url)).then(
        (res) => res.arrayBuffer(),
      );
    const pfpData = await fetch(new URL('https://i.imgur.com/bwzJfrR.jpg', import.meta.url)).then(
        (res) => res.arrayBuffer(),
      );
    return new ImageResponse( 
        (
            <div style={{position: 'relative', display: 'flex'}}>
                 <img 
                        src={imageData}
                    />
                 <div tw="flex flex-col w-full h-full absolute">
                     <div tw="flex h-full w-full">
                         <div tw="flex flex-col md:flex-row w-full py-12 px-4 justify-between p-8">
                            <div tw="flex flex-col text-xl sm:text-3xl mt-6 mr-6 ml-6 mb-6 font-bold tracking-tight text-black text-left">
                                 <span tw="flex flex-col text-lg sm:text-xl md:flex-row w-full py-12 px-4 justify-between p-8" style={{fontFamily: 'Inter-Bold'}}>
                                    <span tw="bottom-16 right-11" style={{fontFamily: 'Inter-Bold', borderRadius: "50%", overflow: 'hidden'}}>
                                        <img tw="w-[180px] h-[182px]"
                                                        src={pfpData}
                                                /> 
                                    </span>
                                 </span>
                            </div>
                         </div>
                     </div>
                 </div>
            </div>
        )
    )
}

