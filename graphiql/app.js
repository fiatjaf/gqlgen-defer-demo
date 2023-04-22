import React from 'react'
import ReactDOM from 'react-dom/client'
import {GraphiQL} from 'graphiql'
import {createClient} from 'graphql-sse'

import 'graphiql/graphiql.css'

ReactDOM.createRoot(document.getElementById('root')).render(
  <GraphiQL fetcher={fetcher} shouldPersistHeaders />
)

let url = '/graphql'

const sse = createClient({
  url,
  singleConnection: false,
  onMessage: ({event, data}) => console.log(event, data),
  retryAttempts: 0,
  headers: () => globalHeaders
})

let globalHeaders = {}

function fetcher({query, operationName, variables}, {headers}) {
  // ugly hack to pass the headers to the sse client
  globalHeaders = headers
  setTimeout(() => {
    globalHeaders = {}
  }, 1)

  let deferred = null
  const pending = []
  let throwMe = null
  let done = false
  const dispose = sse.subscribe(
    {query, operationName, variables},
    {
      next: data => {
        pending.push(data)
        deferred?.resolve(false)
      },
      error: err => {
        throwMe = err
        deferred?.reject(throwMe)
      },
      complete: () => {
        done = true
        deferred?.resolve(true)
      }
    }
  )
  return {
    [Symbol.asyncIterator]() {
      return this
    },
    async next() {
      if (done) return {done: true, value: undefined}
      if (throwMe) throw throwMe
      if (pending.length) return {value: pending.shift()}
      return (await new Promise(
        (resolve, reject) => (deferred = {resolve, reject})
      ))
        ? {done: true, value: undefined}
        : {value: pending.shift()}
    },
    async throw(err) {
      throwMe = err
      deferred?.reject(throwMe)
      return {done: true, value: undefined}
    },
    async return() {
      done = true
      deferred?.resolve(true)
      dispose()
      return {done: true, value: undefined}
    }
  }
}
