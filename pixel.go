package main

import (
    "image"
    "image/color"
    "image/jpeg"
    "image/png"
    "image/gif"
    //"log"
    "os"
    "fmt"
    "bufio"
    "strings"
)
func init() {
    //image.RegisterFormat("jpeg", "\xff\xd8", jpeg.Decode, jpeg.DecodeConfig)
    image.RegisterFormat("jpeg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)
    image.RegisterFormat("gif", "gif", gif.Decode, gif.DecodeConfig)
    image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)

}

func main() {
  
  path := "testdata/shape3.png"
  img, err := imread(path)
  if err != nil {
    fmt.Printf("error loading %s: %v", path,err)
    
  }
  gr := rgbtogray(img)
  out := threshold(gr,200)
  imsave("Grayscale3.jpeg", out)
    
}

func imread(path string) (image.Image,error){
  f, err := os.Open(path)
    if err != nil {
        //log.Fatalf("error loading images.jpg: %v", err)
        return nil,err
    }
    defer f.Close()

    img, _, err := image.Decode(f)
    if err != nil {
        //log.Fatalf("error decoding file to an Image: %v", err)
        return nil, err
    }
    return img, nil
}

func imsave(path string, img image.Image)(bool, error){
  if _, err := os.Stat(path); err != nil && !os.IsNotExist(err) {
        return false,err
    } else if err == nil {
        fmt.Printf("%s already exists. Replace it? (y/n)", path)

        replace, err := bufio.NewReader(os.Stdin).ReadString('\n')
        if err != nil {
            return false,err
            
        } else if strings.TrimRight(strings.ToLower(replace), "\n") != "y" {
          fmt.Printf("%s Creating Overwrite Cancelled successfully", path)
            return false,nil
        }
    }
    
    fmt.Printf("%s Created successfully", path)
    out, err := os.Create(path)
    if err != nil {
      return false,err
    }
    defer out.Close()
    if err := jpeg.Encode(out, img, nil); err != nil {
      return false,err
    }
    return true,nil
}

func threshold (img image.Image, thnum int) image.Image{
      bounds := img.Bounds()
      out := image.NewRGBA(bounds)
    for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
      for x := bounds.Min.X; x < bounds.Max.X; x++{
        r,g,b,_ := img.At(x,y).RGBA()
        r = r>>8
        g = g>>8
        b = b>>8
        //fmt.Println(r,g,b)
        mean := (r+g+b)/3
        if int(mean) > thnum {
          out.Set(x,y, color.RGBA{R:250,G:250,B:250,A:250})
        }else{
          out.Set(x,y, color.RGBA{R:0,G:0,B:0,A:0})
        }
      }
    }
    return out
}

func rgbtogray(img image.Image) image.Image{
	bounds := img.Bounds()
  out := image.NewRGBA(bounds)
    for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
      for x := bounds.Min.X; x < bounds.Max.X; x++{
        r,g,b,_ := img.At(x,y).RGBA()
        r = r>>8
        g = g>>8
        b = b>>8
        f:= 0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b)
        gray := uint8(f + 0.5)
        
        out.Set(x,y, color.RGBA{R:gray,G:gray,B:gray})
      }
    }
    return out
}