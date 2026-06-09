// {{ define "app" }}

const button = document.getElementById("button"),
  password = document.getElementById("password"),
  input = document.getElementById("input");

button.addEventListener("click", async () => {
  try {
    const value = parseInt(input.value);
    if (value < 4) throw "length < 4";
    password.textContent = await generatePassword(value);
  } catch (error) {
    password.textContent = error;
  }
});

// {{ end }}
