'use client'

import { createWeb3Modal, defaultConfig } from '@web3modal/ethers/react'

// 1. Get projectId at https://cloud.walletconnect.com
export const projectId = process.env.NEXT_PUBLIC_PROJECT_ID

let baseRpcUrl = "https://base-mainnet.g.alchemy.com/v2/" + process.env.NEXT_PUBLIC_API_KEY

// 2. Set chains
const mainnet = {
  chainId: 8453,
  name: 'Base Mainnet',
  currency: 'ETH',
  explorerUrl: 'https://basescan.org',
  rpcUrl: baseRpcUrl
}

const metadata = {
    name: 'Versus',
    description: 'Create matches directly on farcaster with Versus, a bot powered by Stadium 🏟️',
    url: 'https://wag3r-bot.vercel.app/', // origin must match your domain & subdomain
    icons: ['https://magenta-hollow-tiglon-795.mypinata.cloud/ipfs/QmWLS8z2ZcQ8JDGLZB9XB4LwXYbnmEQaJiG7fMyvuA42qP']
  }
  
  if (!projectId) {
    throw new Error('Project ID is not defined')
  }

createWeb3Modal({
  ethersConfig: defaultConfig({ metadata }),
  chains: [mainnet],
  projectId,
  enableAnalytics: true // Optional - defaults to your Cloud configuration
})

export function Web3ModalProvider({ children }: any )  {
  return children
}