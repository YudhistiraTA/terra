import LoginForm from "@/components/login/login-form";
import {
  Card,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";

export default function Login() {
  return (
    <div className="flex justify-center items-center min-h-screen p-8 pb-20 gap-16 sm:p-20">
      <Card className="w-[350px] bg-stone-100 shadow-xl shadow-stone-700">
        <CardHeader>
          <CardTitle>Terra</CardTitle>
          <CardDescription>
            <p>Login to your account.</p>
          </CardDescription>
        </CardHeader>
        <LoginForm />
      </Card>
    </div>
  );
}
