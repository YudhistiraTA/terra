import LoginForm from "@/components/login/login-form";
import { Button } from "@/components/ui/button";
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";

export default function Login() {
  return (
    <div className="flex justify-center items-center bg-gray-950 min-h-screen p-8 pb-20 gap-16 sm:p-20">
      <Card className="w-[350px] bg-gray-100 shadow-xl">
        <CardHeader>
          <CardTitle>Terra</CardTitle>
          <CardDescription>
            <p>Login to your account.</p>
          </CardDescription>
        </CardHeader>
        <CardContent>
          <LoginForm />
        </CardContent>
        <CardFooter className="flex justify-between">
          <Button variant="outline">Cancel</Button>
          <Button>Login</Button>
        </CardFooter>
      </Card>
    </div>
  );
}
