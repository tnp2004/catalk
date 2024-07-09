export const CamelToWords = (camel: string): string => {
    return camel.replace(/([a-z])([A-Z])/g, '$1 $2').replace(/^([a-z])/, (match, p1) => p1.toUpperCase())
}