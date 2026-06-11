let LOWERS, UPPERS, DIGITS, OTHERS;
if (typeof charset !== "undefined") {
  LOWERS = charset.lowers;
  UPPERS = charset.uppers;
  DIGITS = charset.digits;
  OTHERS = charset.others;
} else {
  LOWERS = "abcdefghijklmnopqrstuvwxyz";
  UPPERS = LOWERS.toUpperCase();
  DIGITS = "0123456789";
  OTHERS = "`~!@#$%^&*()-_=+[{]}\\|;:'\",<.>/?";
}

function randUint() {
  return Math.floor(Math.random() * 0xff);
}

function randomSource() {
  const fill =
    typeof crypto !== "undefined" && crypto.getRandomValues
      ? (array) => crypto.getRandomValues(array)
      : (array) => array.forEach((_, i, a) => (a[i] = randUint()));

  const CAPACITY = 0x10;
  const array = new Uint8Array(CAPACITY);
  let i = CAPACITY;
  const next = () => (i >= CAPACITY ? (fill(array), (i = 0)) : i++);

  return (source) => source[array[next()] % source.length];
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

const result = passgen(globalThis.LENGTH);

if ("console" in globalThis && console.log) console.log(result);

result;
