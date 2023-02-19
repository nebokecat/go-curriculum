import { gql, useMutation } from '@apollo/client'
import Button from '@mui/material/Button'
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

const IndexPage = () => {
  const [execute] = useMutation(sampleMutation)

  return (
    <Layout title="Home | Next.js + TypeScript Example">
      <h1>Hello Next.js 👋</h1>
      <p>
        <Link href="/about">About</Link>
      </p>
      <Button variant='outlined' onClick={() => execute()}>sample</Button>
    </Layout>
  )
}
export default IndexPage
