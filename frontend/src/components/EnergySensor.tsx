import {Badge} from "@/components/ui/badge"
import {Button} from "@/components/ui/button"
import {Card, CardContent, CardHeader, CardTitle} from "@/components/ui/card"
import {Collapsible, CollapsibleContent, CollapsibleTrigger} from "@/components/ui/collapsible"
import {ChevronDown, ChevronUp, Gauge} from 'lucide-react'
import {useState} from 'react'

interface EnergySensorData {
    apparent_power: number
    available_power: number
    current_summ_delivered: number
    current_tarif: string
    rms_current: number
    rms_current_max: number
}

function SensorMetric(props: { label: string, value: number, unit: string }) {
    return (
        <div className="flex flex-col space-y-1">
            <p className="text-sm font-medium">{props.label}</p>
            <p className="text-2xl font-bold">{props.value} {props.unit}</p>
        </div>
    );
}

export default function EnergySensor({data}: { data: EnergySensorData }) {
    const [isOpen, setIsOpen] = useState(false)

    return (
        <Card className="w-full">
            <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
                <CardTitle className="text-sm font-medium">Energy Consumption</CardTitle>
                <Badge variant="outline">{data.current_tarif}</Badge>
            </CardHeader>
            <CardContent>
                <div className="grid gap-6 sm:grid-cols-2">
                    <div className="flex items-center justify-center p-4 border rounded-lg col-span-full">
                        <Gauge className="h-8 w-8 text-muted-foreground mr-4"/>
                        <div>
                            <p className="text-2xl font-bold">{data.apparent_power} VA</p>
                            <p className="text-xs text-muted-foreground">Apparent Power</p>
                        </div>
                    </div>
                    <SensorMetric label={"Available Power"} value={data.available_power} unit={"A"}/>
                    <SensorMetric label={"Total Energy"} value={data.current_summ_delivered} unit={"kWh"}/>
                    <SensorMetric label={"Current"} value={data.rms_current} unit={"A"}/>
                    <SensorMetric label={"Max Current"} value={data.rms_current_max} unit={"A"}/>
                </div>
            </CardContent>
            <Collapsible open={isOpen} onOpenChange={setIsOpen} className="px-4 pb-4">
                <CollapsibleTrigger asChild>
                    <Button variant="outline" className="w-full justify-between">
                        More
                        {isOpen ? <ChevronUp className="h-4 w-4"/> : <ChevronDown className="h-4 w-4"/>}
                    </Button>
                </CollapsibleTrigger>
                <CollapsibleContent className="mt-4 space-y-4">
                    <div className="grid gap-4 sm:grid-cols-2">
                        <SensorMetricExplanation
                            title="Apparent Power"
                            description="The total power consumed, measured in Volt-Amperes (VA). It's the product of voltage and current."
                        />
                        <SensorMetricExplanation
                            title="Available Power"
                            description="The maximum power that can be drawn from the electrical system, measured in Amperes (A)."
                        />
                        <SensorMetricExplanation
                            title="Total Energy"
                            description="The cumulative energy consumed over time, measured in kilowatt-hours (kWh)."
                        />
                        <SensorMetricExplanation
                            title="Current"
                            description="The amount of electrical current flowing through the system, measured in Amperes (A)."
                        />
                        <SensorMetricExplanation
                            title="Max Current"
                            description="The highest recorded current value, indicating peak usage, measured in Amperes (A)."
                        />
                        <SensorMetricExplanation
                            title="Current Tariff"
                            description="The current pricing scheme for electricity consumption."
                        />
                    </div>
                </CollapsibleContent>
            </Collapsible>
        </Card>
    )
}

function SensorMetricExplanation({title, description}: { title: string, description: string }) {
    return (
        <div className="text-left">
            <h4 className="text-sm font-medium">{title}</h4>
            <p className="text-xs text-muted-foreground">{description}</p>
        </div>
    )
}

