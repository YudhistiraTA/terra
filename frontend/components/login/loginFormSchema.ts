import { z } from "zod";

const loginFormSchema = z.object({
  email: z.string().email(),
  password: z.string().min(6),
});
const loginFormDefaultValue = {
  email: "",
  password: "",
};
export { loginFormSchema, loginFormDefaultValue };
