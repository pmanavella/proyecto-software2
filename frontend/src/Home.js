import React, { useEffect, useState } from 'react';
import './Home.css';
import Cookies from "universal-cookie";
import { FaEmber } from 'react-icons/fa';
import { Link } from 'react-router-dom';

const Cookie = new Cookies();



async function getUserByEmail(email){
    return await fetch('http://localhost:8080/user/' + email, {
    method: 'GET',
    
}).then(response => response.json())
}

// async function getCursoByUserId(userId){
//   return await fetch('http://localhost:8080/course/:id' + userId, {
//     method: "GET",
    
//   }).then(response => response.json())
// }

// courses functions
async function getCourses(){
  return await fetch('http://localhost:8081/course')
    .then(response => response.json())
}

async function SearchByCategory(category){
  return await fetch('http://localhost:8081/course/' + category, {
    method: "GET",
    
  }).then(response => response.json())
}

async function SearchByDescription(description){
  return await fetch('http://localhost:8081/course/' + description, {
    method: "GET",
    
  }).then(response => response.json())
}

async function SearchByTitle(title){
  return await fetch('http://localhost:8081/course/' + title, {
    method: "GET",
    
  }).then(response => response.json())
}

async function postCurso(curso){
  return await fetch('http://localhost:8081/course', { // creatCourse
    method: "POST",
    
    body: JSON.stringify(curso)
  }).then(response => response.json())
}

async function putCurso(curso){
  return await fetch('http://localhost:8081/course/' + curso.id_course, {
    method: "PUT",
    
    body: JSON.stringify(curso)
  }).then(response => response.json())
}

async function creatCourse(curso){
  return await fetch('http://localhost:8081/course', {
    method: "POST",
    
    body: JSON.stringify(curso)
  }).then(response => response.json())
}

async function updateCurso(curso){
  return await fetch('http://localhost:8081/course/' + curso.id_course, {
    method: "PUT",
    
    body: JSON.stringify(curso)
  }).then(response => response.json())
}

async function deleteCurso(curso){
  return await fetch('http://localhost:8081/course/' + curso.id_course, {
    method: "DELETE",
    
  }).then(response => response.json())
}

// extra, search courses about search-api 8082
async function searchCourses(query) {
  try {
    const response = await fetch(`http://localhost:8082/search?query=${query}`, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
      },
    });
    if (!response.ok) {
      throw new Error('Search failed');
    }
    return response.json();
  } catch (error) {
    console.error('Error searching courses:', error);
    throw error;
  }
}

// async function SearchByDescription(description){
//   return await fetch('http://localhost:8081/course/' + description, {
//     method: "GET",
    
//   }).then(response => response.json())
// }

// async function getAvailableCourses(){
//   return await fetch('http://localhost:8081/course/user/available', {
//     method: "POST",
//     body: JSON.stringify({"token":Cookie.get("token")})
//   }).then(response => response.json())
// }

// async function getRegisteredCourses(){
//   return await fetch('http://localhost:8080/course/user/registered', {
//     method: "POST",
//     body: JSON.stringify({"token":Cookie.get("token")})
//   }).then(response => response.json())
// }

async function registerToCourse(id){
  return await fetch("http://localhost:8080/course/register", {
    method: "POST",
    body: JSON.stringify({
      token: Cookie.get("token"),
      course_id: id,
    })
    
  }).then(response => response.json())
}

function goto(path){
  window.location = window.location.origin + path
}


const Home = () => {
  const [admin, setAdmin] = useState(false);
  const [isLogged, setIsLogged] = useState(false);

  const [needCourses, setNeedCourses] = useState(true);
  const [needAvailableCourses, setNeedAvailableCourses] = useState(true);
  const [needRegisteredCourses, setNeedRegisteredCourses] = useState(true);
  
  const [courses, setCourses] = useState([]);
  const [registeredCourses, setRegisteredCourses] = useState([]);
  const [availableCourses, setAvailableCourses] = useState([]);



  
  if(!courses.length && needCourses){
    getCourses().then(response => setCourses(response))
    setNeedCourses(false)
  }
  // if(!availableCourses.length && needAvailableCourses){
  //   getAvailableCourses().then(response => {
  //     if (response) {
  //       setAvailableCourses(response)
  //     }
  //   })
  //   setNeedAvailableCourses(false)
  // }
  // if(!registeredCourses.length && needRegisteredCourses){
  //   getRegisteredCourses().then(response => {
  //     if (response) {
  //       setRegisteredCourses(response)
  //     }
  //   })
  //   setNeedRegisteredCourses(false)
  // }
  

  // const showHomeAdmin = () => {
  //   return (
  //     <div className="container">
  //       <div className="sidebar">
  //         <div className="admin">ADMINISTRADOR</div>
  //         <div className="menu-item">Cursos</div>
  //       </div>
  //       <div className="main-content">
  //         <div className="search-bar">
  //           <input type="text" placeholder="Buscar" />
  //         </div>
  //         <div className="courses">
  //           {courses ? courses.map((course, index) => (
  //             <div key={index} className="Course" onClick={() => goto("/courses/" + course.id_course)}>
  //               <div className="course-item">
  //                 <div>
  //                   <img src={course.image_url} alt={course.title} className="Course-image" />
  //                   <span>{course.title}</span>
  //                 </div>
  //                 <div>
  //                   <p className="course-category">{course.category}</p>
  //                   <p className="course-duration">{course.duration}</p>
  //                   <p className="course-instructor">{course.instructor}</p>
  //                   <p className="course-requirements">{course.requirements}</p>
  //                   <div className="actions">
  //                     <button className="edit">✏️</button>
  //                     <button className="add">+</button>
  //                   </div>
  //                 </div>
  //               </div>
  //             </div>
  //           )) : <p> Loading... </p>}
  //         </div>
  //       </div>
  //       <div className="add-delete-buttons">
  //         <button className="add-course">+</button>
  //         <div className="add-course-text">
  //           <p>+ add new course</p>
  //         </div>
  //         <button className="delete-course">✖</button>
  //         <div className="delete-course-text">
  //           <p>x delete course</p>
  //         </div>
  //       </div>
  //     </div>
  //   );
  // };

  const showHome = () => {
    return (
    <div className="containerAlum">
       {/* <img src='Asserts/Fondo.jpg' alt="Imagen" className="bottom-right-image" /> */}
        <div className="left-section">
          <div className="header">
            <div className="search-bar">
              <input type="text" placeholder="Buscar" />
            </div>
            <Link to="/courses" className="view-search-button">
            🔍
      </Link>
          </div>
          <div className="courses">
          <Link to="/results" className="view-courses-button"> 
        Cursos Disponibles
      </Link>
          <Link to="/courses" className="view-courses-button">
        Mis Cursos
      </Link>
      <Link to="/courses" className="view-courses-button">
        Ver Cursos
      </Link>
            {registeredCourses ? registeredCourses.map((course, index) => (
              <div key={index} className="Course" onClick={() => goto("/courses/" + course.id_course)}>
                <div className="course-item">
                  <div>
                    <img src={course.image_url} alt={course.title} className="Course-image" />
                    <span>{course.title}</span>
                  </div>
                  <div>
                    <p className="course-category">Categoria: {course.category}</p>
                    <p className="course-duration">{course.duration}</p>
                    <p className="course-instructor">{course.instructor}</p>
                    <p className="course-requirements">{course.requirements}</p>
                  </div>
                </div>
              </div>
            )) : <p>Loading courses...</p>}
          </div>
        </div>
        <div className="right-section">
          <div className="available-courses">
          <Link to="/login" className="view-courses-button">
            Log in / Register
      </Link>
            {availableCourses ? availableCourses.map((course, index) => (
              <div key={index} className="Course" onClick={() => goto("/courses/" + course.id_course)}>
                <div className="course-item">
                  <div>
                    <img src={course.image_url} alt={course.title} className="Course-image" />
                    <span>{course.title}</span>
                  </div>
                  <div>
                    <p className="course-duration">{course.duration}</p>
                  </div>
                  <button className="alum-button" onClick={() => registerToCourse(course.id_course)}>INSCRIBIRME</button>
                </div>
              </div>
            )) : <p>Loading Courses...</p>}
          </div>
        </div>
      </div>
    );
  };

  return (
    <div>
      <div className="course-banner">GESTION DE CURSOS</div>
      {showHome()}
    </div>
  );
};

export default Home;
// verificar a donde debe ir el boton login/register