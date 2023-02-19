import { gql } from '@apollo/client'
import Link from 'next/link'
import Layout from '../components/Layout'

// graphql-code-generatorの動作確認用
export const sampleMutation = gql(/* GraphQL */ `
  mutation sample {
    createTodo(input: {userId: "", text: ""}) {
      id
    }
  }
`)

const IndexPage = () => (
  <Layout title="Home | Next.js + TypeScript Example">
    <h1>Hello Next.js 👋</h1>
    <p>
      <Link href="/about">About</Link>
    </p>
  </Layout>
)

export default IndexPage
