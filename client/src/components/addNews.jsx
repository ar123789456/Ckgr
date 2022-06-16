import {React, useFormFields} from 'react'

// const useFormFields = (initialValue) => {
//     const [value, setValue] = React.useState(initialValue);
//     const onChange = React.useCallback((e) => setValue(e.target.value), []);
//     return { value, onChange };
//   };

function AddNews() {
    const { formFields, createChangeHandler } = useFormFields({
        email: '',
        password: '',
      });
    
      const handleSubmit = () => {
        fetch('https://jsonplaceholder.typicode.com/posts', {
            method: 'POST',
            body: JSON.stringify(formFields),
            headers: {
                'Content-type': 'application/json; charset=UTF-8',
            },
        })
        .then((response) => response.json())
        .then((json) => console.log(json));
      };
      return (
        //   <div>Hello</div>
        <form onSubmit={handleSubmit}>
          <div>
            <label htmlFor='email'>Email</label>
            <input type='email' id='email' value={formFields.email} onChange={createChangeHandler('email')} />
          </div>
          <div>
            <label htmlFor='password'>Password</label>
            <input type='password' id='password' value={formFields.password} onChange={createChangeHandler('password')} />
            <input type="submit" />
          </div>
        </form>
      );
}

export default AddNews