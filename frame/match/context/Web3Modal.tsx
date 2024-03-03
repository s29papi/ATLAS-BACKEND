'use client'

import { createWeb3Modal, defaultConfig } from '@web3modal/ethers/react'

// 1. Get projectId at https://cloud.walletconnect.com
export const projectId = process.env.NEXT_PUBLIC_PROJECT_ID

let baseRpcUrl = "https://base-mainnet.g.alchemy.com/v2/" + process.env.NEXT_PUBLIC_API_KEY
let testBaseRpcUrl = "https://base-sepolia.g.alchemy.com/v2/" + process.env.NEXT_PUBLIC_TESTNET_API_KEY
// 2. Set chains
const mainnet = {
  chainId: 8453,
  name: 'Base Mainnet',
  currency: 'ETH',
  explorerUrl: 'https://basescan.org',
  rpcUrl: baseRpcUrl
}

const base_testnet = {
  chainId: 84532,
  name: 'Base Sepolia',
  currency: 'ETH',
  explorerUrl: 'https://sepolia-explorer.base.org',
  rpcUrl: testBaseRpcUrl
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
  chains: [mainnet, base_testnet],
  projectId,
  enableAnalytics: true, // Optional - defaults to your Cloud configuration,
  themeVariables: {
    '--w3m-color-mix': '#00BB7F',
    '--w3m-color-mix-strength': 40
  }
})

export function Web3ModalProvider({ children }: any )  {
  return children
}