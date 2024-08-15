"use client"
import { useState } from 'react';
// import { useRouter } from 'next/router';
import { Conteiner, Container_login} from '@/app/login';

export default function LoginPage() {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState('');
  // const router = useRouter();

  const handleSubmit = async () => {
    // Lógica de autenticación aquí
    try {
      const res = await fetch('/api/auth/login', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ email, password }),
      });

      if (res.ok) {
        // router.push('/dashboard'); // Redirige a la página de dashboard o a la página deseada
      } else {
        const data = await res.json();
        setError(data.message || 'Error en la autenticación');
      }
    } catch (err) {
      setError('Error de red');
    }
  };

  return (
    <Conteiner>
      <Container_login>
        <h1>Login</h1>
        <div className="conteiner-form">
          <form onSubmit={handleSubmit}>
            <div>
              <label>Email:</label>
              <input
                type="email"
                value={email}
                onChange={(e) => setEmail(e.target.value)}
                required
              />
            </div>
            <div>
              <label>Password:</label>
              <input
                type="password"
                value={password}
                onChange={(e) => setPassword(e.target.value)}
                required
              />
            </div>
            <button type="submit">Login</button>
            {error && <p>{error}</p>}
          </form>
        </div>
      </Container_login>
    </Conteiner>
  );
}
