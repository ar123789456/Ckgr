// import logo from './logo.svg';
import Header from './components/header.jsx'
import AddNews from './components/addNews.jsx';
import './App.css';
import {
    BrowserRouter,
    Routes,
    Route,
  } from "react-router-dom";

function App() {
  return (
    <div className="App">
      <BrowserRouter>
      <Header/>

        <Routes>


            <Route path='/' element={<p>Base</p>} />

            <Route path='/projects' element={<p>Projects</p>} />
            <Route path='/projects/:projID' element={<p>single Project</p>} />

            <Route path='/news' element={<p>News</p>} />
            <Route path='/news/:newsID' element={<p>Single News</p>} />

            <Route path='/people' element={<p>People</p>} />

            <Route path='/contact' element={<p>Contact</p>} />

            <Route path='/admin' element={<p>Admin</p>} />
            <Route path='/admin/signin' element={<p>SignIn</p>} />

            <Route path='/admin/add/projects' element={<p>add Project</p>} />
            <Route path='/admin/add/news' element={<AddNews/>} />
            <Route path='/admin/add/user' element={<p>add user</p>} />
            <Route path='/admin/add/contact' element={<p>add contact</p>} />


            <Route path="*" element={
            <main style={{ padding: "1rem" }}>
            <p>404 page not found</p>
            </main>
            }
            />
        </Routes>
      </BrowserRouter>
    </div>
  );
}

export default App;
