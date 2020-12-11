/*
 * author: A. Yasuda
 */

#include <iostream>
#include <string>
#include "libra.h"
#include "crypto.h"
#include "httplib.h"

using namespace httplib;

void agent() {

  std::string op;
  std::cout << " - input hash - \n";
  std::getline(std::cin, op);

  int n = op.length();
  char m[n + 1];
  strcpy(m, op.c_str());
  
  printf("%s, size: %d\n", m, n);
  //GoString tkn = {m, n};
  char *plain = ValidateToken(m);
  printf("return %s, with length %d\n", plain, strlen(plain));
}

void hello() {
    std::cout << "Hello, Libra!" << std::endl;
}

void libra_server() {
  Server svr;

  svr.Get("/index", [](const Request & /*req*/, Response &res) {
		   res.set_content("Hello Libra!", "text/plain");
		 });

  svr.listen("localhost", 8080);
}
