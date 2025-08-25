import React from 'react';
import './home.css';

import gif3 from '../img/gif3.gif'


const Home = () => {

    return (
        <div className="home-container">
            <header className="header">
                <h1> CodeWave Learning</h1>
            </header>

            <div className="content-carousel-container">
                <div className="content-container">
                    <section className="section">
                        <h2 className="section-title">Acerca de Nosotros</h2>
                        <p className="section-content">
                        En CodeWave Learning, nos especializamos en ofrecer una amplia gama de cursos de programación diseñados para personas de todos los niveles, desde principiantes que dan sus primeros pasos en la programación hasta expertos que buscan perfeccionar sus habilidades.
                         Nuestra plataforma está diseñada para proporcionar una experiencia de aprendizaje interactiva y efectiva, centrada en ayudarte a dominar los lenguajes más relevantes en el mundo de la tecnología.
                        </p>
                    </section>
                    <section className="section">
                        <h2 className="section-title">Nuestros Cursos</h2>
                        <p className="section-content">
                            Explora nuestros cursos en diversas tecnologías como JavaScript, Python, Java,
                            y mucho más. Cada curso está diseñado para ofrecer una experiencia de aprendizaje
                            completa y práctica.
                        </p>
                    </section>
                </div>
                <div className="gif-container">  
                <img src={gif3} alt="Descripción del GIF"  className='gif'/>
                </div>

            </div>
            <br/><br/><br/><br/><br/>
            <div className="mision">
                    <div className="mision-container">
                            <section className="section">
                                <h2 className="section-title">Objetivos</h2>
                                <p className="section-content">
                                <br/>
                                • Variedad y Profundidad: Ofrecemos cursos exhaustivos que cubren desde los conceptos básicos hasta temas avanzados en lenguajes como HTML, C++ y Python, así como otras tecnologías y frameworks relevantes en la industria.
                                <br/><br/>
                                • Accesibilidad y Flexibilidad: Nuestros cursos están diseñados para ser accesibles desde cualquier lugar y en cualquier momento, permitiéndote aprender a tu propio ritmo y según tu agenda.
                                <br/><br/>
                                • Enfoque Práctico: Fomentamos un aprendizaje práctico mediante proyectos y ejercicios que te permiten aplicar los conocimientos adquiridos en situaciones reales.
                                <br/><br/>
                                • Apoyo y Comunidad: Estamos comprometidos a ofrecer un soporte integral, tanto a través de nuestros instructores como a través de una comunidad activa de estudiantes que comparten tus intereses y metas.
                                </p>
                            </section>
                         
                     </div>
            </div>


        </div>
    );
};

export default Home;
