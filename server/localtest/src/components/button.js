export const Button = ({name,data,sendData})=>{
    return (
        <>
            <button 
                className='
                    bg-blue-500 hover:bg-blue-700 
                    text-white font-bold py-2 
                    px-4 rounded-full
                    mx-4 my-2
                ' 
                onClick={()=>{sendData(data)}}
            >
                {name}
            </button>
        </>
    )
}