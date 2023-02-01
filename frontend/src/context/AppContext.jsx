import React, { useState, createContext, useEffect } from 'react'
// import { infuraId } from '../utils/app-constants'
import swal from 'sweetalert'

export const AppContext = createContext()

export const AppProvider = ({ children }) => {
  const [loading, setLoading] = useState(false)
  const [currentAccount, setCurrentAccount] = useState()
  const [networkId, setNetworkId] = useState()
  const [networkRpc, setNetworkRpc] = useState()
  const [networkProvider, setNetworkProvider] = useState()
  const [datas, setDatas] = useState([])

  async function insertData(tokenAddress, userAddress) {
    console.log('insertData')
    setLoading(true)
    fetch(
      'http://localhost:8686/insertData?tokenAddress=' +
        tokenAddress +
        '&userAddress=' +
        userAddress,
      {
        mode: 'no-cors',
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
      },
    )
      .then((response) => {
        swal({
          position: 'center',
          icon: 'success',
          title: `Indexes Added Successfully `,
          text: `Just index last 10 txs`,
          button: false,
        })
        getTxDatas(tokenAddress)
        setLoading(false)
      })
      .catch((error) => {
        console.log('Error catch', error)
        swal({
          position: 'center',
          icon: 'error',
          title: `Indexes error`,
          text: `${error}`,
          button: false,
        })
        setLoading(false)
      })
  }

  async function fetchGraphQL(operationsDoc, operationName, variables) {
    const result = await fetch('http://localhost:8080/v1/graphql', {
      method: 'POST',
      body: JSON.stringify({
        query: operationsDoc,
        variables: variables,
        operationName: operationName,
      }),
    })

    return await result.json()
  }

  function fetchMyQuery(tokenAddress) {
    const operationsDoc = `
    query MyQuery {
      transactions(where: {token: {_eq: "${tokenAddress}"}}) {
        _from
        _to
        hash
        input
        success
        timestamp
        value
      }
    }
  `
    console.log('operationsDoc', operationsDoc)
    return fetchGraphQL(operationsDoc, 'MyQuery', {})
  }

  async function getTxDatas(tokenAddress) {
    setLoading(true)
    const { errors, data } = await fetchMyQuery(tokenAddress)
    if (errors) {
      console.error('Hashura error', errors)
    }
    setDatas(data.transactions)
    console.log('Hashura data', data)
  }

  return (
    <AppContext.Provider value={{ insertData, getTxDatas,datas }}>
      {children}
    </AppContext.Provider>
  )
}
