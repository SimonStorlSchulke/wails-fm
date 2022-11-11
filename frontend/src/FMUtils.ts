export class MenuItem {
  public Name: string;
  public Handler: Function;

  constructor(name: string, handler: Function) {
    this.Name = name;
    this.Handler = handler;
  }
}

window.addEventListener("click", (e) => {
  document.querySelector(".rm-menu").classList.add("hidden")
})

function createMenu(items: MenuItem[]): HTMLElement {
  let menu: HTMLElement = document.querySelector(".rm-menu")
  if (!menu) {
    menu = document.createElement("div");
    menu.classList.add("menu")
    menu.classList.add("rm-menu")
  }

  menu.classList.remove("hidden")
  menu.innerHTML = ""

  items.forEach((item) => {
    const b: HTMLElement = document.createElement("button");
    b.addEventListener("click", () => { item.Handler() });
    b.innerHTML = item.Name;
    menu.append(b);
  })

  return menu
}

export function CustomRightclickDialog(element: HTMLElement, items: MenuItem[]) {
  element.addEventListener('contextmenu', function (e: MouseEvent) {
    const menu = createMenu(items);
    document.body.append(menu)
    menu.style.left = (e.pageX - 10) + "px"
    menu.style.top = (e.pageY - 10) + "px"
    e.preventDefault();
  }, false);
}


/**
 * Code by mpen - https://stackoverflow.com/questions/10420352/converting-file-size-in-bytes-to-human-readable-string
 * Format bytes as human-readable text.
 * 
 * @param bytes Number of bytes.
 * @param si True to use metric (SI) units, aka powers of 1000. False to use 
 *           binary (IEC), aka powers of 1024.
 * @param dp Number of decimal places to display.
 * 
 * @return Formatted string.
 */
export function FileSizeText(bytes: number, si = true, dp = 2) {
  const thresh = si ? 1000 : 1024;

  if (Math.abs(bytes) < thresh) {
    return bytes + ' B';
  }

  const units = si
    ? ['kB', 'MB', 'GB', 'TB', 'PB', 'EB', 'ZB', 'YB']
    : ['KiB', 'MiB', 'GiB', 'TiB', 'PiB', 'EiB', 'ZiB', 'YiB'];
  let u = -1;
  const r = 10 ** dp;

  do {
    bytes /= thresh;
    ++u;
  } while (Math.round(Math.abs(bytes) * r) / r >= thresh && u < units.length - 1);


  return bytes.toFixed(dp) + ' ' + units[u];
}