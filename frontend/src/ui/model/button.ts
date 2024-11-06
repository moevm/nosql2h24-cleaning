type Button = {
  id?: string,
  class?: string,
  color?: string | undefined,
  variant?: "text" | "flat" | "elevated" | "tonal" | "outlined" | "plain" | undefined
  isActive?: boolean,
  type?: string,
  text: string,
  to?: string,
  onClick?: Function
}

export default Button;