document.addEventListener('DOMContentLoaded', () => {
  // --- Copy Code Functionality ---
  const copyButtons = document.querySelectorAll('.copy-code-button');

  copyButtons.forEach(button => {
    button.addEventListener('click', () => {
      const wrapper = button.parentElement;
      const codeElement = wrapper.querySelector('code');
      
      if (!codeElement) return;

      let codeText = codeElement.innerText;

      navigator.clipboard.writeText(codeText).then(() => {
        button.classList.add('copied');
        const copyIcon = button.querySelector('.copy-icon');
        const checkIcon = button.querySelector('.check-icon');
        
        if (copyIcon && checkIcon) {
          copyIcon.style.display = 'none';
          checkIcon.style.display = 'block';
          
          setTimeout(() => {
            copyIcon.style.display = 'block';
            checkIcon.style.display = 'none';
            button.classList.remove('copied');
          }, 2000);
        }
      }).catch(err => {
        console.error('Failed to copy text: ', err);
      });
    });
  });
});
