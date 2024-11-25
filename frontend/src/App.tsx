import './App.css'
import EnergySensor from "@/components/EnergySensor.tsx"
import Logo from "@/components/Logo.tsx";
import {RefreshIntervalSwitcher, RefreshOption} from "@/components/RefreshInterval.tsx"
import {SiGithub} from "@icons-pack/react-simple-icons";
import {useQuery} from "@tanstack/react-query"
import {BarChart2, Code, Heart,} from 'lucide-react'
import path from "path-browserify"

import React from "react"

const refreshOptions: RefreshOption[] = [
    {label: 'Off', value: "0"},
    {label: '1 second', value: "1000"},
    {label: '5 seconds', value: "5000"},
    {label: '10 seconds', value: "10000"},
    {label: '30 seconds', value: "30000"},
    {label: '1 minute', value: "60000"},
    {label: '5 minutes', value: "300000"},
    {label: '15 minutes', value: "900000"},
    {label: '30 minutes', value: "1800000"},
    {label: '1 hour', value: "3600000"},
    {label: '2 hours', value: "7200000"},
    {label: '1 day', value: "86400000"},
]

function App() {
    const [refreshInterval, setRefreshInterval] = useLocalStorage("refresh-interval", refreshOptions[0].value)

    const {isPending, error, data} = useQuery({
        refetchInterval: parseInt(refreshInterval),
        queryKey: ['LiXee'], queryFn: async () => {
            const response = await fetch(`${path.join(import.meta.env.BASE_URL, import.meta.env.VITE_API_ENDPOINT || "api/v1/lixee")}`)
            return await response.json()
        },
    })

    if (isPending) return 'Loading...'

    if (error) return 'An error has occurred: ' + error.message

    return (
        <div className="flex flex-col items-center justify-center min-h-screen bg-background p-4">
            <a href="https://github.com/merlindorin" className="flex items-center" target="_blank"
               rel="noopener noreferrer">
                <Logo width={64} fill="#EEE" className="mb-12"/>
            </a>
            <div style={{minWidth: 580}} className="w-full max-w-xl mx-auto px-4 sm:px-6 lg:px-8">
                <div className="flex justify-between items-center mb-4">
                    <h1 className="text-2xl font-bold">Lixee - Linky</h1>
                    <div className="flex items-center space-x-2">
                        <label htmlFor="auto-refresh" className="text-sm text-muted-foreground mr-2">
                            Auto-refresh:
                        </label>
                        <RefreshIntervalSwitcher
                            refreshOptions={refreshOptions}
                            onValueChange={setRefreshInterval}
                            currentValue={refreshInterval}
                        />
                    </div>
                </div>
                <EnergySensor data={data}/>
                <footer className="mt-8 pb-4 flex justify-center items-center space-x-4 text-sm text-muted-foreground">
                    <a href="/metrics" className="flex items-center hover:text-foreground">
                        <BarChart2 className="w-4 h-4 mr-1"/>
                        Metrics
                    </a>
                    <span>/</span>
                    <a href="/-/healthy" className="flex items-center hover:text-foreground">
                        <Heart className="w-4 h-4 mr-1"/>
                        Health
                    </a>
                    <span>/</span>
                    <a href="/api/v1/lixee" className="flex items-center hover:text-foreground">
                        <Code className="w-4 h-4 mr-1"/>
                        API
                    </a>
                    <span>/</span>
                    <a
                        href="https://github.com/merlindorin/exporter-lixee"
                        className="flex items-center hover:text-foreground"
                        target="_blank"
                        rel="noopener noreferrer"
                    >
                        <SiGithub className="w-4 h-4 mr-1"/>
                        Github
                    </a>
                </footer>
            </div>

        </div>
    )
}

const useLocalStorage = (keyName: string, defaultValue: any) => {
    const [storedValue, setStoredValue] = React.useState(() => {
        try {
            const value = window.localStorage.getItem(keyName);

            if (value) {
                return JSON.parse(value);
            } else {
                window.localStorage.setItem(keyName, JSON.stringify(defaultValue));
                return defaultValue;
            }
        } catch (err) {
            return defaultValue;
        }
    });

    const setValue: (v: string) => void = newValue => {
        try {
            window.localStorage.setItem(keyName, JSON.stringify(newValue));
        } catch (err) {
        }
        setStoredValue(newValue);
    };

    return [storedValue, setValue];
};
export default App
