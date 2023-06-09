import React, { useEffect, useState } from 'react';

const Blicky = ({type, className}) => {

  const [vis, setVis] = useState(false)

  useEffect(() => {
    setInterval(() => {
      setVis(old => !old)
    }, 700);
  }, []);

  return React.createElement(type, { className: className, style : {color : "transparent", backgroundColor : vis ? "black" : "transparent"} }, <>i</>);
}

export default function TypeWriterWrapper({type, className, content = "EMPTY", delay = 100, onAnimationEnd = () => {}}) {

  const [finished, setFinished] = useState(false)

  return (
    <div style={{display : "flex", alignItems : "center"}}>
      <TypeWriter type={type} className={className} content={content} onAnimationEnd={setFinished} />
      {
        !finished &&
        <Blicky type={type} />
      }

    </div>
  )

}

function TypeWriter({type, className, content, delay = 100, onAnimationEnd = () => {}}) {
  const [text, setText] = useState('');
  const message = content;


  useEffect(() => {
    let i = 0;
    const type = setInterval(() => {
      setText(message.slice(0, i));
      i++;
      if (i > message.length) {
        clearInterval(type);
        onAnimationEnd(true)
      }
    }, delay);
  }, []);

  return React.createElement(type, { className: className }, <>{text}</>);
}


