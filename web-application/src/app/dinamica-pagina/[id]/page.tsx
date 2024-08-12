/**
 * As paginas dinamica segueem o principio das paginas esttias porem com um pequeno detanhe onde
 * dentro do arquivo principal deve ter um outra pasta nomeada dentro de cochete[] onde dentro desse cochete
 * voce deve fazer referencia ao que vai ser passado na URL como Id por exemplo.
 */

/**
 * Essa diretiva 'fala' para a aplicação que ao rodoar essa apagina ela vai precisar de 
 * elementos de renderização que estaram na pagina de navegação do usuario.
 */
"use client"

// Biblioteca para capturar valores de Url
import { useParams } from "next/navigation";

export default function PageStatic() {
    // Iniciando a variavel que vai armazenaar o valor do paramentro da URL
    const parametro_id = useParams();

    return (
      <main className="flex min-h-screen flex-col items-center justify-between p-24">
        <div className="z-10 w-full max-w-5xl items-center justify-between font-mono text-sm lg:flex">
         pagina dinamica 
         valor de URL {parametro_id.id}
        </div>
      </main>
    );
  }
  