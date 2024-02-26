/* eslint-disable @next/next/no-img-element  */
/* eslint-disable jsx-ally/alt-text */
// @ts-nocheck
import { ImageResponse } from 'next/server'

export const runtime = 'edge';


export async function GET(req: Request) {
    const imageData = await fetch(new URL('../../../public/You-Vs-Me-Rescale.png', import.meta.url)).then(
        (res) => res.arrayBuffer(),
      );
    const interData = await fetch(new URL('../../../public/Inter-Regular.ttf', import.meta.url)).then(
        (res) => res.arrayBuffer(),
      );
    const { searchParams } = new URL(request.url);
    const gameName = searchParams.get('gameName');
    console.log(gameName)
    let stakeAmount = "25 USDC"
    let gameName = "Tekken 8"
    let gameSetup = "Best of 4"
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
                                                    src={'https://i.imgur.com/bwzJfrR.jpg'}
                                            /> 
                                        
                                        </span>  
                                        <span tw="flex flex-col bottom-7" style={{fontFamily: 'Inter-Regular'}}>
                                           <span tw="text-5xl text-gray-400">{stakeAmount} open challenge</span>
                                           <span tw="text-7xl top-[0.95]">{gameName}{" "}/{" "}{gameSetup}</span>
                                        </span>  
                                    </span>  
                                
                                </div>
                            </div>
                        </div>
                </div>
        </div>
        ),
            {
                fonts: [
                    {
                        name: "Inter-Regular",
                        data: interData,
                        style: "normal"
                    }
                ]
            }
    )
}

