import {Select, SelectContent, SelectItem, SelectTrigger, SelectValue,} from "@/components/ui/select"

import {cn} from "@/lib/utils"
import {type SelectTriggerProps} from "@radix-ui/react-select"


export interface RefreshOption {
    label: string
    value: string
}

interface RefreshIntervalSwitcherProps extends SelectTriggerProps {
    onValueChange(value: string): void
    refreshOptions: RefreshOption[]

    currentValue: string
}

export function RefreshIntervalSwitcher({className, onValueChange, currentValue, refreshOptions, ...props}: RefreshIntervalSwitcherProps) {
    return (
        <Select
            value={currentValue}
            onValueChange={onValueChange}
        >
            <SelectTrigger
                className={cn(
                    "h-7 w-[145px] text-xs [&_svg]:h-4 [&_svg]:w-4",
                    className
                )}
                {...props}
            >
                <SelectValue placeholder="Select Refresh Interval"/>
            </SelectTrigger>
            <SelectContent>
                {refreshOptions.map((opts) => (
                    <SelectItem  key={opts.label} value={opts.value} className="text-xs">
                        {opts.label}
                    </SelectItem>
                ))}
            </SelectContent>
        </Select>
    )
}