import React, { useState } from 'react';
import SignInSignUp from "./pages/SignInSignUp"

export default function App() {

  const [user, setUser] = useState(null);


  return (
    <div>
      {user ? (
      <h1>Estas logueado</h1>
      ) : (
      <div>
        <SignInSignUp />
      </div>
      )}
    </div>
  )
}