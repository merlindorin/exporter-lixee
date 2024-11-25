import {SVGProps} from "react"

const SvgComponent = (props: SVGProps<SVGSVGElement>) => (
    <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 640 640" {...props}>
        <path
            d="m134.9 556.09 355-236.7-170.26-113.52L134.9 329Zm405.54-236.7a18 18 0 0 1-8 15L126.86 604.76a18 18 0 0 1-28-15V319.39a18 18 0 0 1 8-15l202.78-135.18a18 18 0 0 1 20 0L532.4 304.39a18 18 0 0 1 8 15"/>
        <path
            d="m319.64 432.9 170.27-113.51-355-236.7v227.05Zm220.8-113.51a18 18 0 0 1-8 15L329.63 469.57a18.07 18.07 0 0 1-20 0L106.87 334.38a18 18 0 0 1-8-15V49a18 18 0 0 1 28-15L532.4 304.39a18 18 0 0 1 8 15"/>
    </svg>
)
export default SvgComponent
