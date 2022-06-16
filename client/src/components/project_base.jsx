import React from 'react'
// import Project_block from './project_block'

export default function ProjectBase() {
    // let [news, setNews] = useState()
    
    return (
        <div className='projects-par'>
            <div className='project-header'>
                <div>
                    <h1>Projects</h1>
                    <p>dddddddddddddddddddddddddddddd</p>
                </div>
                
                <iframe title='1' width="420" height="315" src="https://www.youtube.com/embed/tgbNymZ7vqY">
                </iframe>            
            </div>
            <div className='project-body'>
                {/* <Project_block />
                <Project_block />
                <Project_block />
                <Project_block /> */}
            </div>
        </div>
    )
    // useEffect(()=> {
    //     fetch('https://jsonplaceholder.typicode.com/todos')
    //     .then(response => response.json())
    //     .then(json => setNews(json))
    // })
    // return (
    //     <div>
    //         <h1>News</h1>
    //         <pre>{JSON.stringify(news, null, 2)}</pre>

    //     </div>
    // )
}
