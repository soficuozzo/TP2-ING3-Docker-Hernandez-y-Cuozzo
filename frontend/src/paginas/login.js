import React, { useState } from 'react';
import { useNavigate } from "react-router-dom";
import toast, { Toaster } from 'react-hot-toast';
import './login.css';
import { Link } from 'react-router-dom';

const LoginForm = () => {
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");
    const navigate = useNavigate(); // Permite la navegación entre páginas con las rutas

    const handleSubmit = async (e) => {
        e.preventDefault();
        
        // Validación básica del lado del cliente (puedes personalizar según tus necesidades)
        if (!email.trim()) {
            toast.error("Ingrese el email correctamente.");
            return;
        }

        if (!password.trim()) {
            toast.error("Ingrese la contraseña correctamente.");
            return;
        }
        
        try {
            const response = await fetch('http://localhost:8081/login', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({ email, password }),
            });

            if (response.ok) {
                const data = await response.json();
                localStorage.setItem('token', data.token); // Almacena el token en localStorage
                
                // Obtener datos del usuario y almacenar en localStorage
                const userData = await fetchUserData(email);
                if (userData) {
                    localStorage.setItem("email", userData.email);
                    localStorage.setItem("tipo", userData.tipo);
                    localStorage.setItem("usuario_id", userData.usuario_id);
                }
                toast.success("Inicio de sesión exitoso");
            navigate('/');
               
            }  else {
                    const errorMessage = await response.json(); // Leer el mensaje de error del servidor
                    if (response.status === 401 && errorMessage.message === "Login no autorizado: Hubo un error al buscar el usuario en la Base de Datos.") {
                        toast.error("El email ingresado no existe. Por favor, verifique sus credenciales.");
                    } else {
                        toast.error("Error al intentar iniciar sesión");
                    }
                }
        } catch (error) {
            console.error('Error al realizar la solicitud al backend:', error);
            toast.error("Error al intentar iniciar sesión");
        }
    };

    const fetchUserData = async (email) => {
        try {
            const response = await fetch(`http://localhost:8081/users/email/${email}`, {
                headers: {
                    Authorization: `Bearer ${localStorage.getItem('token')}`, // Agrega el token de autenticación en los headers
                },
            });
            if (response.ok) {
                return await response.json(); // Devuelve los datos del usuario si la solicitud es exitosa
            } else {
                console.error('Error al obtener datos del usuario:', response.status);
                return null; // Devuelve null si la solicitud no es exitosa
            }
        } catch (error) {
            console.error('Error al realizar la solicitud al backend:', error);
            return null; // Devuelve null si hay un error durante la solicitud
        }
    };

    return (
        <div className="login-form-container">
            <form onSubmit={handleSubmit} className="login-form">
                <div className="form-group">
                    <label>Email</label>
                    <br/>
                    <input
                        id="inputEmailLogin"
                        type="text"
                        value={email}
                        onChange={(e) => setEmail(e.target.value)}
                        placeholder="Ingrese su email"
                        required
                    />
                </div>
                <div className="form-group">
                    <label>Contraseña</label>
                    <input
                        id="inputPasswordLogin"
                        type="password"
                        value={password}
                        onChange={(e) => setPassword(e.target.value)}
                        placeholder="Ingrese su contraseña"
                        required
                    />
                </div>
                <br/><br/>
                <button type="submit">Iniciar sesión</button>
                <br/><br/>
                <Link to={`/registrarse`} className="curso-link">
                                            <strong>No tenes cuenta? Registrate</strong>
                                        </Link>
            </form>


            <Toaster /> {/* Componente para mostrar notificaciones toast */}
        </div>
    );
};

export default LoginForm;
