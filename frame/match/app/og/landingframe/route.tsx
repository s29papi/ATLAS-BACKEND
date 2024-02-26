/* eslint-disable @next/next/no-img-element  */
/* eslint-disable jsx-ally/alt-text */
// @ts-nocheck
import { ImageResponse } from 'next/server'

export const runtime = 'edge';


export async function GET(req: Request) {

    return new ImageResponse( 
        (
            <div>Hello eooieo kjjdjdjdj</div>
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

