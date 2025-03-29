import { useState } from 'react'
// import './App.css'
import { postLogin, postRegister } from './api';
import { initializeBg } from './bg';

const App = () => {
    initializeBg('bg');

    const [user, setUser] = useState("");
    const [pass, setPass] = useState("");

    // todo write (simple) reactquery yourself basically

    return (
        <>
            <canvas id='bg'></canvas>

            <h1>
                <div><a href="http://monke.ca">monke.ca</a></div>
                <div>
                    storage pleasure
                </div>
            </h1>
            <form className='form-section'>
                <input
                    aria-label='username'
                    type='text'
                    id='username'
                    name='username'
                    onChange={(v) => setUser(v.target.value || "")}
                    value={user}
                    placeholder='username'
                />

                <input
                    // aria-details='password-rules'
                    aria-label='password'
                    type='password'
                    onChange={(v) => setPass(v.target.value || "")}
                    value={pass}
                    placeholder='password'
                />

                {/*
                <p id='password-rules'>
                    passwords may be from 5 to 70 characters long, (inclusive)

                </p>*/}

                <div className="form-buttons">
                    <button type="submit" aria-label='login' onClick={() => postLogin(user, pass)} >
                        login
                    </button>
                    <button type="submit" aria-label='register' onClick={() => postRegister(user, pass)} >
                        register
                    </button>
                </div>
            </form>
        </>
    )
}

export default App
