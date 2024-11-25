import {ReactQueryDevtools} from "@tanstack/react-query-devtools";
import {StrictMode} from 'react'
import {createRoot} from 'react-dom/client'

import {QueryClient, QueryClientProvider} from "@tanstack/react-query";

import './index.css'
import App from './App.tsx'

const queryClient = new QueryClient()

createRoot(document.getElementById('root')!).render(
    <StrictMode>
        <QueryClientProvider client={queryClient}>
            <ReactQueryDevtools initialIsOpen={false} />
            <App/>
        </QueryClientProvider>
    </StrictMode>,
)