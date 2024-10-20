import { SignedOut, SignOutButton, UserButton } from "@clerk/nextjs"

export default function Home() {
  return (
    <div>
      <UserButton
                appearance={{
                    elements: {
                        userButtonPopoverFooter: {
                            signOutButton: {
                                signOutUrl: "/"
                            }
                        }
                    }
                }}
            />
    </div>
  )
}