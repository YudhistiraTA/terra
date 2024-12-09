import { z } from "zod";

const postFormSchema = z.object({
  title: z.string().min(1),
  content: z.string().min(1),
});
const postFormDefaultValue = {
  title: "",
  content: "",
};
export { postFormSchema, postFormDefaultValue };
