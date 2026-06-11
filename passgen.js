function randomSource() {
  if (globalThis.crypto) {
    const CAPACITY = 0x10;

    const refresh = () => {
      crypto.getRandomValues(array);
      i = 0;
    };

    let array = new Uint8Array(CAPACITY),
      i = CAPACITY;

    return (source) => {
      if (i >= CAPACITY) refresh();
      return source[array[i++] % source.length];
    };
  } else {
    return (source) => source[Math.floor(Math.random() * source.length)];
  }
}

function shuffle(string) {
  return [...string].sort(() => Math.random() - 0.5).join("");
}

function passgen(length = 8) {
  if (length < 4) throw new Error("length < 4");

  const randomChar = randomSource();

  let password = "";
  password += randomChar(LOWERS);
  password += randomChar(UPPERS);
  password += randomChar(DIGITS);
  password += randomChar(OTHERS);

  const CHARSET = LOWERS + UPPERS + DIGITS + OTHERS;

  for (let i = 4; i < length; i++) password += randomChar(CHARSET);

  return shuffle(password);
}

passgen(LENGTH);
