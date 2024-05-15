import { Alert, AlertDescription, AlertTitle } from '@/components/ui/alert'
import { Input } from '@/components/ui/input'
import { createLazyFileRoute } from '@tanstack/react-router'
import { Coffee } from 'lucide-react'

export const Route = createLazyFileRoute('/')({
  component: Index,
})

function Index() {
  return <div className="p-2 container">
    <Alert>
      <Coffee />
      <AlertTitle>Early Development!</AlertTitle>
      <AlertDescription>
        SmartWiki is currently in early development and not yet usable. I appreciate your interest and encourage you to check back later for updates as we work towards a functional release
      </AlertDescription>
    </Alert>

    <div className="text-center py-8">
      <h1 className='text-4xl'>SmartWiki</h1>
      <h2 className='text-2xl'>Your Knowledge, Simplified and Amplified</h2>
      <Input className='mt-4' placeholder='Search for Documentation...' />
    </div>
  </div>

}
