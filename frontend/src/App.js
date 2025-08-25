import './App.css';
import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
import NavegadorHome from './componentes/NavegadorHome';
import Cursos from './paginas/cursos';
import Curso from './paginas/curso';
import Home from './paginas/home';
import Resultados from './paginas/resultados';
import LoginForm from './paginas/login';
import Registrarse from './paginas/registrarse';
import MisCursos from './paginas/miscursos';
import CerrarSesion from './paginas/cerrarsesion';
import CursoNuevo from './paginas/cursonuevo';
import UpdateCurso from './paginas/updatecurso';
import Contenedores from './paginas/contenedores';


import { useState } from 'react';

function App() {
  const [searchTerm, setSearchTerm] = useState('');

  return (
    <div>
      <Router>
        <NavegadorHome onSearch={setSearchTerm} />
        <Routes>
          <Route exact path="/" element={<Home />} /> 
          <Route exact path="/cursos/todos" element={<Cursos />} /> 
          <Route exact path="/curso/:curso_id" element={<Curso />} /> 
          <Route exact path="/cursos/buscar" element={<Resultados searchTerm={searchTerm}/>} /> 
          <Route exact path="/login" element={<LoginForm />} /> 
          <Route exact path="/registrarse" element={<Registrarse />} /> 
          <Route exact path="/miscursos" element={<MisCursos />} /> 
          <Route exact path="/cerrarsesion" element={<CerrarSesion />} /> 
          <Route exact path="/cursonuevo" element={<CursoNuevo />} /> 
          <Route exact path="/updatecurso/:curso_id" element={<UpdateCurso />} /> 
          <Route exact path="/servicios" element={<Contenedores />} /> 


        </Routes>
      </Router>

      <footer />
    </div>
  );
}

export default App;
