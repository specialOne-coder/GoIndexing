import React, { useContext, useState } from 'react'
import { Loader, Welcome } from '../components/index'
import { FcMoneyTransfer } from 'react-icons/fc'
import { AppContext } from '../context/AppContext'
import { ethereumClient, wagmiClient } from '../App'
import { Web3Button } from '@web3modal/react'
import { useEffect } from 'react'
import { shortenAddress } from '../utils/ShortAddress'

const companyCommonStyles =
  'min-h-[70px] sm:px-0 px-2 sm:min-w-[220px] flex justify-center items-center border-[0.5px] border-gray-400 text-sm font-light text-white'

const IndexesPage = () => {
  const { insertData, getTxDatas, datas } = useContext(AppContext)
  const [isCon, setCon] = useState(false)
  const [tokenAddr, setTokenAddr] = useState(false)
  const [fileName, setFileName] = useState('')
  console.log('Hashura in indexes', datas)
  let dataLength = datas.length
  let composeFileName;

  useEffect(() => {
    const interval = setInterval(async () => {
      const isConnected = ethereumClient.getAccount().isConnected
      if (isConnected) {
        setCon(true)
      } else {
        setCon(false)
      }
    }, 1000)
    return () => clearInterval(interval)
  }, [])

  function timestampToDate(timestamp) {
    var date = new Date(timestamp * 1000)
    return date.toDateString()
  }
  return (
    <div>
      <center>
        <div className="flex flex-row justify-center items-center ">
          <div className="flex-initial w-64  border-white rounded-lg text-white mt-2 py-2 px-1 ">
            <input
              type="text"
              id="tokenAddress"
              className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-[15rem] p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500 mt-1"
              placeholder="your token address"
              min="1"
              required
            />
          </div>
          <div className=" flex-initial w-32  border-white ml-2 rounded-lg text-white mt-1 px-2 py-1">
            {isCon ? (
              <button
                onClick={async () => {
                  const tokenAddress = document.getElementById('tokenAddress')
                    .value
                  const userAddress = ethereumClient.getAccount().address
                  let file = 'txs_'+tokenAddress+'_'+userAddress+'.json'
                  setFileName(file)
                  if (tokenAddress !== '' && userAddress !== '') {
                    await insertData(tokenAddress, userAddress)
                    await getTxDatas(tokenAddress)
                  }
                }}
                className="flex flex-row justify-center w-full items-center bg-gradient-to-r from-cyan-500 to-blue-500 p-2 rounded-lg cursor-pointer mt-2"
              >
                <p className="text-white text-base font-semibold flex flex-row">
                  indexe
                </p>
              </button>
            ) : (
              <Web3Button icon="show" label="Connect Wallet" />
            )}
          </div>
          <div className=" flex-initial w-32  border-white ml-2 rounded-lg text-white mt-1 px-2 py-1">
            {dataLength > 0 ? (
            <button
                className="flex flex-row justify-center w-full items-center bg-gradient-to-r from-cyan-500 to-blue-500 p-2 rounded-lg cursor-pointer mt-2"

            onClick={() => {
              fetch(
                `http://localhost:8686/download?filename=${fileName}`,
              ).then((res) => {
                res.blob().then((blob) => {
                  let url = window.URL.createObjectURL(blob)
                  let a = document.createElement('a')
                  a.style.display = 'none'
                  a.href = url
                  a.download = fileName
                  document.body.appendChild(a)
                  a.click()
                  window.URL.revokeObjectURL(url)
                })
              })
              console.log('file name', fileName);
            }}
            >
              {} Download
            </button>
            ) : (
              ''
            )}
          </div>
        </div>
      </center>
      <div className="flex w-full justify-center items-center">
        <div className="flex mf:flex-row flex-col items-center justify-between  py-12 px-4">
          <div className="flex flex-1 justify-center items-center flex-col  mf:ml-30">
            <h1 className="text-3xl text-white  py-1"> Last 10 txs </h1>
            <div className="grid sm:grid-cols-6 grid-cols-2 w-full mt-5">
              <div
                className={`rounded-tl-2xl sm:rounded-bl-2xl ${companyCommonStyles}`}
              >
                <p className="font-bold text-[20px]">Hash</p>
              </div>
              <div className={companyCommonStyles}>
                <p className="font-bold text-[20px]">Time</p>
              </div>
              <div className={companyCommonStyles}>
                <p className="font-bold text-[20px]">From</p>
              </div>
              <div className={companyCommonStyles}>
                <p className="font-bold text-[20px]">To</p>
              </div>
              <div className={companyCommonStyles}>
                <p className="font-bold text-[20px]">Value</p>
              </div>
              <div
                className={`sm:rounded-tr-2xl rounded-br-2xl ${companyCommonStyles}`}
              >
                <p className="font-bold text-[20px]">Succes</p>
              </div>
            </div>
            {dataLength == 0 ? (
              <>
                <h1 className="text-3xl text-white  py-5">Loading... </h1>
                <Loader />
              </>
            ) : (
              [...datas].map((data, i) => (
                <div className="grid sm:grid-cols-6 w-full" key={i}>
                  <>
                    <div
                      className={`rounded-tl-2xl sm:rounded-bl-2xl ${companyCommonStyles}`}
                    >
                      <p className="">{shortenAddress(data.hash)}</p>
                    </div>

                    <div className={companyCommonStyles}>
                      {timestampToDate(data.timestamp)}
                    </div>
                    <div className={companyCommonStyles}>
                      {shortenAddress(data._from)}
                    </div>
                    <div className={companyCommonStyles}>
                      {shortenAddress(data._to)}
                    </div>
                    <div className={companyCommonStyles}>{data.value}</div>
                    <div
                      className={`sm:rounded-tr-2xl rounded-br-2xl ${companyCommonStyles}`}
                    >
                      {data.success}
                    </div>
                  </>
                </div>
              ))
            )}
          </div>
        </div>
      </div>
    </div>
  )
}

export default IndexesPage
